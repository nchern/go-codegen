## impl command

Generates stubs to implement a given interface
The difference between this command and impl(https://github.com/josharian/impl) utility is that interface declaration is read from stdin.
So that it's really easy to use it with editors like Vim

### Command line

```
$ echo "type Connection interface { Open() }" |go-codegen impl

type connection struct {}

func (c *connection) Open() {
	panic("Not implemented")
}
```

### Vim integration

It could be more convenient to use this generator from your favourite text editor. Here is an example how to integrate it with Vim.
Just add the following line to your `~/.vimrc`:

```vim
autocmd FileType go command! -range=% GoGenImpl :<line1>,<line2>!go-codegen impl -s
```

Now you can select the struct in VISUAL mode and run the command `:GoGenImpl` to generate constructor for it.
