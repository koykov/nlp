package nlp

//go:generate nlpc --module=languages --input=raw/languages1.json --target=languages_repo.go
//go:generate nlpc --module=scripts --input=raw/scripts.json --target=scripts_repo.go

//go:generate nlpc -module=ngmodels -input="https://github.com/koykov/dataset/blob/master/vocabulary/freelang/English.txt?raw=true" -target=~/dev/dataset/ngmodels/English.ngm
//go:generate nlpc -module=ngmodels -input="https://github.com/koykov/dataset/blob/master/vocabulary/freelang/Russian.txt?raw=true" -target=~/dev/dataset/ngmodels/Russian.ngm
