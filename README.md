# devotion-project-template

A template that can be used to "aldev bootstrap" a new [Devotion](https://aldesgroup.github.io/devotion/) project.

It can contain:

- an API part, based on `Goald` (Go)
- and / or a web app, based on `GoaldR` (React)
- and / or a native app, based on `GoaldN` (React Native)

## How to use:

Make sure you have done the prerequisites steps of [Devotion](https://aldesgroup.github.io/devotion/).

Then:

```sh
cd $GITPRIV &&                                 \
aldev bootstrap --private --web-less --verbose \
--name FooBar                                  \
--repo my-git.my-domain.com                    \
--group group_name[/subgroup] &&               \
cd foo-bar
```

This creates a new project `FooBar` then git-sync it with `my-git.my-domain.com/group_name[/subgroup]/foo-bar`.

After that, you're ready to [start working with your new project](https://aldesgroup.github.io/devotion/docs/start/bootstrap#working-with-the-new-api).
