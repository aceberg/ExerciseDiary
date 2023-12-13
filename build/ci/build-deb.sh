#!/bin/bash

PKGNAME=AppTemplate
PKGDESC="Write app description here"
PKGDIR=$PKGNAME-$1-0_all

umask 0022

mkdir -p $PKGDIR/usr/bin
mkdir -p $PKGDIR/lib/systemd/system

cp configs/$PKGNAME.service $PKGDIR/lib/systemd/system/
cp configs/$PKGNAME@.service $PKGDIR/lib/systemd/system/

cp $PKGNAME $PKGDIR/usr/bin/

mkdir -p $PKGDIR/DEBIAN

echo "Package: $PKGNAME
Version: $1
Section: utils
Priority: optional
Architecture: all
Maintainer: aceberg <aceberg_a@proton.me>
Description: $PKGDESC
" > $PKGDIR/DEBIAN/control

echo "
systemctl daemon-reload
" > $PKGDIR/DEBIAN/postinst

chmod 775 $PKGDIR/DEBIAN/postinst

dpkg-deb --build --root-owner-group $PKGDIR

rm -rf $PKGDIR