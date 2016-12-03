#! /bin/sh
npm install --silent
npm run build
# ln -s dist/* /srv/
npm run watch
