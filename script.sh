IMG=qemu-image.img
DIR=temp.dir

sudo apt-get install qemu-system debootstrap

qemu-img create $IMG 1g
mkfs.ext2 $IMG
mkdir $DIR
sudo mount -o loop $IMG $DIR

echo "Starting bootstrap..." 

sudo debootstrap --arch amd64 stable $DIR

echo "Configuring autologin..."
sudo mkdir $DIR/etc/systemd/system/getty@tty1.service.d
sudo tee $DIR/etc/systemd/system/getty@tty1.service.d/autologin.conf > /dev/null <<EOF
[Service]
ExecStart=
ExecStart=-/sbin/agetty -o '-p -f -- \\u' --noclear --autologin root %I $TERM
EOF

echo "Replacing motd"
sudo rm $DIR/etc/motd

echo "Putting 'hello world' into motd so it gets printed on boot."
sudo tee $DIR/etc/motd > /dev/null <<EOF
hello world
EOF

echo "Unmounting"
sudo umount $DIR

echo "Deleting mount point dir"
rmdir $DIR

echo "Launching qemu"
sudo qemu-system-amd64 -kernel /boot/vmlinuz-$(uname -r) -drive file=$IMG,index=0,media=disk,format=raw -append "root=/dev/sda"
