# Tips & Tricks

## Workflow
Personally, I use [`entr`](https://github.com/eradman/entr) to automatically run `mdtmpl` every time my template file (e.g `README.md.tmpl`) changes. Simply output the **filenames** you want to watch on `STDOUT` and pipe them into `entr` with the command to be executed at file change:

```bash
> echo README.md.tmpl | entr mdtmpl -f
```
