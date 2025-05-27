{
  inputs = {
    nixos-generators = {
      url = "github:nix-community/nixos-generators";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    };
  };
  outputs =
    { nixpkgs, ... }@inputs:
    let
      each =
        function:
        nixpkgs.lib.genAttrs [
          "x86_64-darwin"
          "x86_64-linux"
          "aarch64-darwin"
          "aarch64-linux"
        ] (system: function nixpkgs.legacyPackages.${system});
    in
    {
      devShells = each (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            awscli2
            gnumake
            go
            gocyclo
            gofumpt
            golangci-lint
            gopls
            ntp
            opentofu
          ];
          shellHook = ''
            go mod tidy
          '';
        };
      });
      packages.x86_64-linux = {
        tullius =
          let
            cicero = nixpkgs.legacyPackages."x86_64-linux".buildGoModule {
              pname = "cicero";
              version = "now";
              src = ./cicero;
              ldflags = [
                "-s"
                "-w"
              ];
              doCheck = false;
              vendorHash = "sha256-yXXLs0NV7jQhRMCyWy8wbYQGRJXv8RLHFIYZI1EryWM=";
            };
            configurations = {
              system.stateVersion = "24.05";
              networking = {
                hostName = "tullius";
                firewall = {
                  enable = true;
                  allowedTCPPorts = [ 80 ];
                  allowedUDPPorts = [ 123 ];
                };
                networkmanager = {
                  enable = true;
                };
              };
              systemd.services.cicero = {
                enable = true;
                wantedBy = [ "multi-user.target" ];
                after = [ "network.target" ];
                script = ''
                  ${cicero}/bin/cicero
                '';
                serviceConfig = {
                  Restart = "always";
                  Type = "simple";
                };
              };
              time = {
                timeZone = "Etc/UTC";
              };
            };
            name = "cicero-x86_64-linux";
            image = inputs.nixos-generators.nixosGenerate {
              format = "amazon";
              modules = [
                configurations
                {
                  amazonImage = {
                    inherit name;
                    sizeMB = 4 * 1024;
                  };
                }
              ];
            };
            virtualization = "${image}/${name}.vhd";
          in
          {
            inherit cicero;
            tofu = nixpkgs.legacyPackages."x86_64-linux".writeShellScriptBin "tofu" ''
              export TF_VAR_image="${virtualization}"
              ${nixpkgs.legacyPackages."x86_64-linux".opentofu}/bin/tofu $@
            '';
          };
      };
    };
}
