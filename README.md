# docker-composer

`docker-composer` is a tool for mixing and matching common elements of dockerfiles.

## Usage

The program can be run without arguments for an interactive mode, or with arguments:

### Build a dockerfile from a template

```bash
docker-composer dockerfile "template_name" "mixin_name_1" "mixin_name_2" ... "mixin_name_n"
```

### Manage template

```bash
docker-composer template "create"/"edit"/"delete" "template_name"
```

### Manage mixin

```bash
docker-composer mixin "create"/"edit"/"delete"  "mixin_name"
```
