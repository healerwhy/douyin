package globalkey

// 存放的是api记录点赞和关注放在redis的key与valTpl

const FavoriteSetKey = "FavoriteKey"
const FavoriteSetValTpl = "Favorite&Video&Id:%d"

const FollowSetKey = "FollowKey"
const FollowSetValTpl = "Follow&Follow&Id:%d"

const ExistDataValTpl = "%d:%d"
