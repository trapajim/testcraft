package datagen

type dictionary struct {
	Adjective []string
	Verbs     []string
	Nouns     []string
	Domains   []string
}

var dict *dictionary

func Dict() *dictionary {
	if dict == nil {
		dict = &dictionary{Adjective: adjectives, Verbs: verbs, Nouns: nouns, Domains: domains}
	}
	return dict
}

func (d *dictionary) RandomAdjective() string {
	return d.Adjective[Rand().Int(len(d.Adjective))]
}

func (d *dictionary) RandomVerb() string {
	return verbs[Rand().Int(len(verbs))]
}

func (d *dictionary) RandomNoun() string {
	return nouns[Rand().Int(len(nouns))]
}

func (d *dictionary) RandomDomain() string {
	return domains[Rand().Int(len(domains))]
}
