{
  description = "setups to manage setups";
  inputs = {
    flake-utils = {
      url = "github:numtide/flake-utils";
    };
    nixos-generators = {
      url = "github:nix-community/nixos-generators";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    nixpkgs = {
      url = "github:NixOS/nixpkgs";
    };
  };
  outputs = { nixpkgs, ... }@inputs:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs {
        inherit system;
      };
      cicero = pkgs.buildGoModule {
        pname = "cicero";
        version = "now";
        src = ../cicero;
        ldflags = [ "-s" "-w" ];
        doCheck = false;
        vendorHash = "sha256-yXXLs0NV7jQhRMCyWy8wbYQGRJXv8RLHFIYZI1EryWM=";
      };
      configurations = {
        system.stateVersion = "24.05";
        networking = {
          firewall = {
            allowedTCPPorts = [ 443 ];
            allowedUDPPorts = [ 123 ];
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
              sizeMB = 4 * 1024;
            };
          }
        ];
      };
      virtualization = "${image}/${name}.vhd";
    in
    {
      devShell.${system} = pkgs.mkShell {
        buildInputs = [
          pkgs.awscli2
          pkgs.opentofu
        ];
      };
      packages.${system} = {
        inherit cicero;
        tofu = pkgs.writeShellScriptBin "tofu" ''
          export TF_VAR_image="${virtualization}"
          ${pkgs.opentofu}/bin/tofu $@
        '';
      };
    };
}
