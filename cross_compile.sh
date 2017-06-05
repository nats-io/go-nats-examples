#!/bin/bash
go get github.com/mitchellh/gox
go get github.com/tcnksm/ghr

export OUTDIR="release"

# If we have an arg, assume its a version tag and rename as appropriate.
if [[ -n $1 ]]; then
    export APPNAME=$APPNAME-$1
fi

env CGO_ENABLED=0 gox -arch amd64 -output "release/{{.OS}}_{{.Arch}}/{{.Dir}}" $(go list ./... | grep -v "/vendor/")

for dir in $(ls release); do \
    (cp LICENSE $OUTDIR/$dir/LICENSE) ;\
    (cd $OUTDIR && zip -q go-nats-examples-$dir.zip -r $dir) ;\
    echo "make $OUTDIR/go-nats-examples-$dir.zip" ;\
done

