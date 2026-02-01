{
  inputs = {
    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    };
  };
  outputs =
    { nixpkgs, ... }:
    let
      each =
        function:
        nixpkgs.lib.genAttrs [
          "aarch64-darwin"
          "aarch64-linux"
          "x86_64-darwin"
          "x86_64-linux"
        ] (system: function nixpkgs.legacyPackages.${system});
    in
    {
      devShells = each (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            git # https://github.com/git/git
            gnumake # https://github.com/mirror/make
            go # https://github.com/golang/go
            gocyclo # https://github.com/fzipp/gocyclo
            gofumpt # https://github.com/mvdan/gofumpt
            golangci-lint # https://github.com/golangci/golangci-lint
            gopls # https://github.com/golang/tools/tree/master/gopls
            ntp # https://www.ntp.org/documentation/4.2.8-series/
            tailwindcss_4 # https://github.com/tailwindlabs/tailwindcss
          ];
          shellHook = ''
            go mod tidy
          '';
        };
      });
      packages = each (pkgs: {
        default = pkgs.buildGoModule {
          pname = "cicero";
          version = "now";
          src = ./.;
          ldflags = [
            "-s"
            "-w"
          ];
          doCheck = true;
          vendorHash = "sha256-yXXLs0NV7jQhRMCyWy8wbYQGRJXv8RLHFIYZI1EryWM=";
        };
      });
    };
}
