# IDE Untuk Go

IDE (`Integrated Development Environment`) merupakan software yang digunakan oleh pemrogram dan pengembang software untuk membangun software. IDE berisi berbagai fasilitas komprehensif yang diperlukan para pemrogram untuk membantu mereka dalam membangun software aplikasi. Secara minimal, biasanya IDE terdiri atas editor teks (untuk mengetikkan kode sumber), debugger (pencari bugs), `syntax highlighting, code completion, serta dokumentasi / help. Bab ini akan membahas beberapa software yang bisa digunakan. Sebenarnya menggunakan editor teks yang menghasilkan file text / ASCII murni sudah cukup untuk bisa menuliskan dan kemudian mengkompilasi kode sumber. Pada bab ini akan dibahas [Vim](http://www.vim.org) sebagai editor teks dan LiteIDE sebagai software IDE yang lebih lengkap untuk Go, tidak sekedar hanya untuk menuliskan kode sumber.

## Menggunakan Vim

Untuk menggunakan Vim, ada plugin utama serta berbagai plugin pendukung yang bisa digunakan. Sebaiknya, menggunakan `pathogen` untuk mempermudah pengelolaan berbagai plugin tersebut. Bagian ini akan menjelaskan berbagai konfigurasi serta instalasi yang diperlukan sehingga Vim bisa menjadi peranti untuk pengembangan aplikasi menggunakan Go.

## Instalasi dan Konfigurasi Pathogen

Pathogen adalah plugin dari Tim Pope yang digunakan untuk mempermudah pengelolaan plugin. Kode sumber dari Pathogen bisa diperoleh di [repository Github](https://github.com/tpope/vim-pathogen). Untuk instalasi, ikuti langkah berikut:

~~~bash
$ cd
$ mkdir .vim/autoload
$ mkdir .vim/bundle
$ cd .vim/autoload
$ wget https://raw.githubusercontent.com/tpope/vim-pathogen/master/autoload/pathogen.vim
--2016-08-15 11:16:50--  https://raw.githubusercontent.com/tpope/vim-pathogen/master/autoload/pathogen.vim
Resolving raw.githubusercontent.com (raw.githubusercontent.com)... 151.101.100.133
Connecting to raw.githubusercontent.com (raw.githubusercontent.com)|151.101.100.133|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 12474 (12K) [text/plain]
Saving to: ‘pathogen.vim’

pathogen.vim                   100%[==================================================>]  12.18K  --.-KB/s    in 0.02s   

2016-08-15 11:16:56 (512 KB/s) - ‘pathogen.vim’ saved [12474/12474]

$
~~~

Setelah itu, untuk menggunakan Pathogen, letakkan aktivasinya di `$HOME/.vimrc` atau di `$HOME/.vim/vimrc` (saya pilih lokasi yang kedua) sebagai berikut:

~~~bash
execute pathogen#infect()
~~~

Setelah itu, semua plugin tinggal kita ambil dari repository (bisa dari github, bitbucket, dan lain-lain) langsung di-copy satu direktori ke direktori `$HOME/.vim/bundle`.

##  Instalasi dan Kofigurasi Plugin Golang dan Plugin Pendukung

Setelah selesai melakukan instalasi Pathogen, berbagai plugin yang diperlukan bisa diambil langsung dari Internet. Berikut ini adalah daftar yang digunakan penulis:

* [Colorschemes](https://github.com/flazz/vim-colorschemes): untuk tema warna dari Vim.
* [Nerdtree](https://github.com/scrooloose/nerdtree): untuk menampilkan file-file dalam struktur pohon di sebelah kiri sehingga memudahkan navigasi.
* [vim-go](https://github.com/fatih/vim-go.git): plugin utama agar Vim mengenali kode sumber Go.

Cara instalasi:

~~~bash
$ cd 
$ cd .vim/bundle
$ git clone <masing-masing lokasi plugin>
~~~

Hasil dari menjalankan ``vim'' atau ``gvim'' melalui shell untuk menulis kode sumber Go bisa dilihat pada gambar berikut ini:

![vim-go](images/vim-go.jpg)

## Autocompletion

Vim menyediakan fasilitas `autocompletion` melalui `Omniautocompletion`. Fasilitas ini sudah terinstall saat kita menginstall Vim. Untuk mengaktifkan fasilitas ini untuk keperluan Go, kita harus menginstall dan mengaktifkan [Gocode](https://github.com/nsf/gocode). Gocode sudah terinstall setelah selesai mengerjakan instalasi di [bab 1](bab-01.md). Hasil dari instalasi Gocode adalah file `executable binary` `$GOPATH/bin/gocode` (sesuai letak GOPATH di env.sh). Sebelum menggunakan Vim, aktifkan dulu gocode dengan mengeksekusi `gocode` melalui shell. Setelah itu, tambahkan satu baris di `$HOME/.vim/vimrc`: `set ofu=syntaxcomplete#Complete` di bawah baris `filetype plugin indent on`.

Kode sumber lengkap dari `$HOME/.vim/vimrc` yang penulis gunakan bisa dilihat pada listing berikut ini:

~~~vim
execute pathogen#infect()

syntax on
filetype plugin indent on
set ofu=syntaxcomplete#Complete

if has("gui_running")
	colorscheme asmanian_blood
else
	colorscheme slate
endif

set smartindent
set tabstop=2
set shiftwidth=2
set expandtab

autocmd vimenter * NERDTree
autocmd vimenter * if !argc() | NERDTree | endif
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTreeType") && b:NERDTreeType == "primary") | q | endif

let g:NERDTreeDirArrows=0

let g:cssColorVimDoNotMessMyUpdatetime = 1
set guifont=Liberation\ Mono\ 11

set number
set numberwidth=4
set cpoptions+=n
highlight LineNr term=bold cterm=NONE ctermfg=DarkGrey ctermbg=NONE gui=NONE guifg=DarkGrey guibg=NONE

set grepprg=grep\ -nH\ $*
~~~

Untuk mengaktifkan completion, kita harus masuk ke mode `Insert` dari Vim, setelah itu tekan `Ctrl-X, Ctrl-O` secara cepat. Hasil `autocompletion` bisa dilihat di gambar berikut ini:

![Go completion](images/vim-go-completion.jpg)

## Menggunakan LiteIDE

![LiteIDE](images/liteide.png)

LiteIDE dibuat oleh visualfc dan tersedia dalam bentuk kode sumber maupun binary. Kode sumber bisa diperoleh di [repo GitHub](https://github.com/visualfc/liteide). Installer executable bisa diperoleh di [Sourceforge](http://sourceforge.net/projects/liteide/files)

Instalasi di Linux sangat mudah, hanya tinggal mengekstrak file yang kita download pada suatu direktori dan jika ingin menjalankan cukup dengan mengeksekusi file `$LITEIDE_HOME/bin/liteide` (`cd $LITEIDE_HOME/bin; ./liteide &`)

## Software IDE Lain

Vim dan LiteIDE hanyalah beberapa peranti yang bisa digunakan oleh pengembang. Distribusi Go juga menyediakan dukungan untuk berbagai peranti lunak lain:

* Emacs. Dukungan untuk Go diwujudkan dalam fasilitas `add-on`. Untuk Emacs 24 ke atas, bisa diinstall melalui manajer paket (M-x package-list-packages), cari dan install `go-mode`. Emacs juga mendukung `gocode` untuk `completion`.
* Eclipse. Dukungan untuk Go diwujudkan melalui plugin `goclipse`, bisa diperoleh di https://code.google.com/p/goclipse/.
* Selain software-software yang telah disebutkan, rata-rata IDE / Editor sudah mempunyai dukungan terhadap bahasa pemrograman Go (JEdit, Sublime-text, Notepad++, dan lain-lain).

