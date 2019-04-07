with import <nixpkgs>{};
buildGoPackage rec {
  name = "formats-${version}";
  version = "development";

  buildInputs = with pkgs; [ git dep ];
  installPhase = ''
    source $stdenv/setup
    set -e

    mkdir -p             $bin/bin
    cp    go/bin/formats $bin/bin
  '';

  src = ./.;
  goDeps = ./deps.nix;
  goPackagePath = "github.com/corpix/formats";
}
