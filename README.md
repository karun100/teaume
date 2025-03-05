# teaume--command to generate ent code automtically schema file
kparuthi@kparuthi-mbp tea_ume % go generate ./ent
kparuthi@kparuthi-mbp tea_ume % 

# teaume--command to restore generate file
kparuthi@kparuthi-mbp tea_ume % git restore ent/generate.go
kparuthi@kparuthi-mbp tea_ume % 
# teaume--regenerate
go run entgo.io/ent/cmd/ent generate ./ent/schema
