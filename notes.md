## References

https://github.com/ikegami-yukino/mecab-as-kkc
http://www.tahoo.org/~taku/software/mecab-skkserv/
https://qiita.com/nownabe/items/4171776aec1f05de9f28
https://github.com/aankittcoolest/mecab-as-kkc
https://qiita.com/yukinoi/items/f136c0cbf3b0e07b175e
http://www.tahoo.org/~taku/software/mecab-skkserv/
https://taku910.github.io/mecab/dic.html
https://qiita.com/YuukiMiyoshi/items/00b9878a1fa32b859a43
http://www.ajisaba.net/server/mecab.html
http://www.chasen.org/~taku/software/mecab-skkserv/

https://google.com/complete/search?output=toolbar&gl=jp&q=%E3%82%B5%E3%82%AB%E3%82%A4
https://github.com/ts-3156/validate_japanese
https://github.com/tomoemon/text_normalizer
https://qiita.com/akifumii/items/bf1511cb8bc53e12f503
https://github.com/ikawaha/kagome


## Try mecab-skkserv

```sh
cd /tmp
wget http://www.tahoo.org/~taku/software/mecab-skkserv/mecab-skkserv-0.03.tar.gz
tar -xvf mecab-skkserv-0.03.tar.gz
cd mecab-skkserv-0.03/
apk add alpine-sdk
apk add gcc autoconf automake
```

## Install mecab and mozc

```sh
apk add --update build-base
cd /tmp
wget 'https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7cENtOXlicTFaRUE' -O mecab-0.996.tar.gz
tar -xvf mecab-0.996.tar.gz
cd mecab-0.996/
./configure  --with-charset=utf8 --enable-utf8-only --build=aarch64-unknown-linux-gnu
make
make install

cd /tmp
wget 'https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM' -O mecab-ipadic-2.7.0-20070801.tar.gz
tar zxvf mecab-ipadic-2.7.0-20070801.tar.gz
cd mecab-ipadic-2.7.0-20070801
./configure  --with-charset=utf8
make
make install

./configure --with-mecab-config=/opt/mecab/bin/mecab-config --with-charset=utf8
./configure --with-charset=utf8
make
sudo make install


cd /tmp
apk add git python3
git clone --depth 1 https://github.com/ikegami-yukino/mecab-as-kkc.git
cd mecab-as-kkc/
make
make install
```