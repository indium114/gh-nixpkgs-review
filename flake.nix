{
  description = "go devshell and package, created by scaffolder";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "go-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
            just
          ];
        };

        packages.example = pkgs.buildGoModule {
          pname = "gh-nixpkgs-review";
          version = "0.1.0";

          src = self;

          vendorHash = null;

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A very small gh extension to provide a more convenient way to use nixpkgs-review-gha";
            license = licenses.mit;
            platforms = platforms.all;
          };
        };

        apps.example = {
          type = "app";
          program = "${self.packages.${pkgs.stdenv.hostPlatform.system}.example}/bin/gh-nixpkgs-review";
        };
      });
}
