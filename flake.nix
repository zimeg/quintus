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
            system = "x86_64-linux";
            pkgs = import nixpkgs {
              inherit system;
            };
            cicero = pkgs.buildGoModule {
              pname = "cicero";
              version = "now";
              src = ./cicero;
              ldflags = [
                "-s"
                "-w"
              ];
              doCheck = true;
              vendorHash = "sha256-yXXLs0NV7jQhRMCyWy8wbYQGRJXv8RLHFIYZI1EryWM=";
            };
            configurations = {
              system.stateVersion = "24.05";
              nix.registry = {
                nixpkgs.flake = nixpkgs;
              };
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
              virtualisation = {
                diskSize = 4 * 1024;
              };
            };
            name = "cicero-${system}";
            image = inputs.nixos-generators.nixosGenerate {
              inherit pkgs;
              format = "amazon";
              modules = [
                configurations
                {
                  amazonImage = {
                    inherit name;
                  };
                }
              ];
            };
            virtualization = "${image}/${name}.vhd";
          in
          {
            inherit cicero;
            tofu = pkgs.writeShellScriptBin "tofu" ''
              export TF_VAR_image="${virtualization}"
              ${pkgs.opentofu}/bin/tofu $@
            '';
          };
      };
    };
}
