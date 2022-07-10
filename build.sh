#!/bin/bash

go build .
rm -rf ./build
mkdir build
cp -r views ./build/
cp -r static ./build/
mv ./snowcraft-website ./build/
sed -i 's/<script src=\"\/static\/tailwind.js\"><\/script>/<link rel=\"stylesheet\" href=\"\/static\/tailwind.min.css\">/g' ./build/views/layouts/master.html
rm ./build/static/tailwind.js
tailwind -i ./tailwind.css -o ./build/static/tailwind.min.css --minify

