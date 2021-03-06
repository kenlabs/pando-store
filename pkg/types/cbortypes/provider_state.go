package cbortypes

//go:generate cbor-gen-for ProviderState

//type ProviderState struct {
//	MetaStateMap map[string]*MetaState
//}

// ProviderStateRes for graphql,
// include the total state about the provider and the newest state change(such as cidlist)
//type ProviderStateRes struct {
//	State        ProviderState
//	NewestUpdate []cid.Cid
//}

//func (t *ProviderState) String() string {
//	str := ""
//	str += fmt.Sprintf("last Commit height: %d, ", t.LastCommitHeight)
//	str += "cidlist: ["
//	for _, c := range t.MetaList {
//		str += fmt.Sprintf("%s, ", c.String())
//	}
//	str += "]"
//	return str
//}
