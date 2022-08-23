# Ringkasan Materi #02 - Fundamental Git

Git adalah salah satu version control system yang wajib diketahui oleh programmer terutama ketika bekerja dalam tim. Git berada pada proyek perangkat lunak yang diciptakan oleh Linus Torvalds.


## Git Command
Ada beberapa general command dari git, yaitu :

|command|fungsi|
|:------------- | :----------|
|`git init <options>`| membuat git repository baru|
| `git add <file>` atau `git add <directory>` | menambahkan file atau direktori pada repositori |
| `git remote <options>` | |
| `git status` | menampilkan keadaan working directory dan staging area |
| `git commit <options>` | menangkap perubahan proyek yang saat ini |
| `git push <options>` | mengunggah commit ke remote repository |
| `git fetch <remote> [options]` | mengambil info meta-data terbaru dari yang asli ke git lokal |
| `git pull <remote> <branch_name>` | mendapatkan pembaruan dari remote repo |
| `git branch <branch-name>` | membuat cabang secara lokal |
| `git checkout <name-branch>` | berpindah dari satu branch ke branch lainnya |
| `git merge <name-branch>` | menggabungkan beberapa urutan commit menjadi satu |
| `git stash <command>` | menyimpan perubahan |


## Commit Convention
Commit convention merupakan salah satu cara untuk membawakan productive testing environment kepada developer karena akan memudahkan dalam development ke depannya. Selain itu, juga berfungsi untuk :
- automatic generating pada changelog
- navigasi sederhana melalui git history

format :

`<type>(<scope>): <subject>`

`<BLANK LINE>`

`<body>`

`<BLANK LINE>`

`<footer>`


Allowed `<type>`values:

- **feat** for a new feature for the user, not a new feature for build script. Such commit will trigger a release bumping a MINOR version.
- **fix** for a bug fix for the user, not a fix to a build script. Such commit will trigger a release bumping a PATCH version.
- **perf** for performance improvements. Such commit will trigger a release bumping a PATCH version.
- **docs** for changes to the documentation.
- **style** for formatting changes, missing semicolons, etc.
- **refactor** for refactoring production code, e.g. renaming a variable.
- **test** for adding missing tests, refactoring tests; no production code change.
- **build** for updating build configuration, development tools or other changes irrelevant to the user.


## Semantic Versioning

`v<major>.<minor>.<patch>`

Example: v2.0.0

**Patch** : fix, perf

**Minor** : feat

**Major** : breaking changes in the API


## Git Management
Di eFishery git management yang digunakan adalah Trunk Based Developement.

Trunk Based Development menerapkan : 
- berkolaborasi pada kode dalam satu cabang yang disebut 'trunk' (utama atau master)
- menghindari long-lived branches
- Gunakanshort-lived branches hanya untuk pengembangan fitur
- Gabungkan pull request hanya jika sudah di-review
- Masing-masing dan setiap commit yang dibuat ke trunk dianggap sebagai unit yang dapat digunakan