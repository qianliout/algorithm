" ------------------------------
" 插件管理器配置
" ------------------------------
call plug#begin('~/.vim/plugged')

" 插件列表
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }  " Go 语言支持
Plug 'neoclide/coc.nvim', {'branch': 'release'}     " LSP 和自动补全
Plug 'junegunn/fzf', { 'do': { -> fzf#install() } } " fzf 依赖
Plug 'junegunn/fzf.vim'                             " 快速搜索
Plug 'jiangmiao/auto-pairs'                         " 自动补全括号
Plug 'vim-airline/vim-airline'                      " 状态栏
Plug 'vim-airline/vim-airline-themes'               " 状态栏主题
Plug 'dracula/vim', { 'as': 'dracula' }             " Dracula 主题
Plug 'preservim/nerdtree'                           " 文件树插件
Plug 'tpope/vim-commentary'                         " 注释插件
Plug 'stephpy/vim-yaml'                            " YAML 支持
Plug 'elzr/vim-json'                                " JSON 支持
Plug 'plasticboy/vim-markdown'                      " Markdown 支持
Plug 'iamcco/markdown-preview.nvim', { 'do': 'cd app && yarn install' }  " Markdown 实时预览
Plug 'lervag/vimtex'                                " LaTeX 支持
Plug 'xuhdev/vim-latex-live-preview'                " LaTeX 实时预览
Plug 'ekalinin/Dockerfile.vim'                      " Dockerfile 支持
Plug 'davidhalter/jedi-vim'                         " Python 支持
Plug 'vim-python/python-syntax'                     " Python 语法高亮
Plug 'psf/black', { 'branch': 'stable' }            " Python 格式化工具
Plug 'neoclide/coc-python', {'do': 'yarn install --frozen-lockfile'}  " Python LSP 支持
Plug 'puremourning/vimspector'                      " Python 调试支持
Plug 'dhruvasagar/vim-table-mode'                   " Markdown 表格支持

call plug#end()

" ------------------------------
" 基础配置
" ------------------------------
set number                      " 显示行号
set relativenumber              " 显示相对行号
set tabstop=4                   " Tab 键的宽度
set shiftwidth=4                " 自动缩进的宽度
set expandtab                   " 将 Tab 转换为空格
set autoindent                  " 自动缩进
set smartindent                 " 智能缩进
set cursorline                  " 高亮当前行
set showmatch                   " 显示匹配的括号
set ignorecase                  " 搜索时忽略大小写
set smartcase                   " 搜索时智能大小写
set hlsearch                    " 高亮搜索结果
set incsearch                   " 增量搜索
set mouse=a                     " 启用鼠标支持
set encoding=utf-8              " 设置编码
set fileencoding=utf-8          " 设置文件编码
set nobackup                    " 不创建备份文件
set nowritebackup               " 不创建写入备份文件
set noswapfile                  " 不创建交换文件
set clipboard=unnamed           " 使用系统剪贴板

" 设置 Leader 键为空格键
let mapleader = " "

" ------------------------------
" 主题配置
" ------------------------------
colorscheme dracula             " 使用 Dracula 主题
syntax enable                   " 启用语法高亮
set termguicolors               " 启用真彩色支持
let g:airline_theme='dracula'   " 状态栏主题

" ------------------------------
" vim-go 配置
" ------------------------------
let g:go_fmt_command = "goimports"  " 使用 goimports 进行格式化
let g:go_auto_type_info = 1         " 自动显示类型信息
let g:go_highlight_functions = 1    " 高亮函数
let g:go_highlight_methods = 1      " 高亮方法
let g:go_highlight_structs = 1      " 高亮结构体
let g:go_highlight_interfaces = 1   " 高亮接口
let g:go_highlight_operators = 1    " 高亮操作符
let g:go_def_mode = 'gopls'         " 使用 gopls 进行定义跳转
let g:go_info_mode = 'gopls'        " 使用 gopls 获取类型信息

" ------------------------------
" coc.nvim 配置
" ------------------------------
let g:coc_global_extensions = ['coc-go', 'coc-json', 'coc-yaml', 'coc-python']  " 启用语言支持

" 使用 gopls 作为 LSP 服务器
let g:coc_config = {
  \ 'languageserver': {
  \   'golang': {
  \     'command': 'gopls',
  \     'rootPatterns': ['go.mod', '.git/'],
  \     'filetypes': ['go'],
  \     'initializationOptions': {
  \       'usePlaceholders': v:true,
  \       'completeUnimported': v:true,
  \       'staticcheck': v:true,
  \     },
  \   },
  \   'yaml': {
  \     'command': 'yaml-language-server',
  \     'filetypes': ['yaml'],
  \     'args': ['--stdio'],
  \     'settings': {
  \       'yaml': {
  \         'schemas': {
  \           'kubernetes': '/*.yaml',  " 支持 Kubernetes YAML 文件
  \         },
  \       },
  \     },
  \   },
  \   'json': {
  \     'command': 'vscode-json-language-server',
  \     'filetypes': ['json'],
  \     'args': ['--stdio'],
  \     'settings': {
  \       'json': {
  \         'schemas': [
  \           {
  \             'fileMatch': ['package.json'],
  \             'url': 'https://json.schemastore.org/package.json',
  \           },
  \           {
  \             'fileMatch': ['tsconfig.json'],
  \             'url': 'https://json.schemastore.org/tsconfig.json',
  \           },
  \         ],
  \       },
  \     },
  \   },
  \   'python': {
  \     'command': 'pylsp',
  \     'filetypes': ['python'],
  \     'settings': {
  \       'pylsp': {
  \         'plugins': {
  \           'pycodestyle': {
  \             'enabled': v:true,
  \           },
  \           'pylint': {
  \             'enabled': v:true,
  \           },
  \           'black': {
  \             'enabled': v:true,
  \           },
  \         },
  \       },
  \     },
  \   },
  \ },
\ }

" ------------------------------
" 自定义函数：根据文件类型格式化
" ------------------------------
function! FormatFile()
    let l:filetype = &filetype  " 获取当前文件的文件类型

    if l:filetype == 'python'
        :Black  " 使用 Black 格式化 Python 文件
    elseif l:filetype == 'go'
        :call CocAction('format')  " 使用 Coc 格式化 Go 文件
    elseif l:filetype == 'json'
        :%!jq .  " 使用 jq 格式化 JSON 文件
    elseif l:filetype == 'yaml'
        :%!yq eval .  " 使用 yq 格式化 YAML 文件
    elseif l:filetype == 'markdown'
        :%!markdownfmt  " 使用 markdownfmt 格式化 Markdown 文件
    elseif l:filetype == 'dockerfile'
        :%!dockerfile-format  " 使用 dockerfile-format 格式化 Dockerfile
    elseif l:filetype == 'tex'
        :%!latexindent  " 使用 latexindent 格式化 LaTeX 文件
    else
        echo "Unsupported file type for formatting: " . l:filetype
    endif
endfunction

" ------------------------------
" 快捷键配置
" ------------------------------
" 按 leader+ff 根据文件类型格式化当前文件
nnoremap <leader>ff :call FormatFile()<CR>

" 按 leader+n 打开/关闭 NERDTree
nnoremap <leader>n :NERDTreeToggle<CR>

" 按 leader+f 打开文件搜索
nnoremap <leader>f :Files<CR>

" 按 leader+b 打开缓冲区搜索
nnoremap <leader>b :Buffers<CR>

" 按 leader+g 打开内容搜索
nnoremap <leader>g :Rg<CR>

" 按 leader+cr 重命名符号
nmap <leader>cr <Plug>(coc-rename)

" 按 leader+mp 打开/关闭 Markdown 预览
nnoremap <leader>mp :MarkdownPreviewToggle<CR>

" 按 leader+lp 打开/关闭 LaTeX 实时预览
nnoremap <leader>lp :LLPStartPreview<CR>

" 按 leader+tm 切换表格模式
nnoremap <leader>tm :TableModeToggle<CR>
" 按 leader+w 保存文件
nnoremap <leader>w :w<CR>

" 按 leader+s 保存并退出
nnoremap <leader>sq :wq<CR>