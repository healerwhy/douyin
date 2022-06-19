package globalkey

/**
redis key except "model cache key"  in here,
but "model cache key" in model
*/

//// CacheUserTokenKey /** 用户登陆的token
//const CacheUserTokenKey = "user_token:%d"

const FavoriteSetKey = "FavoriteKey"
const FavoriteTpl = "Favorite&Video&Id:%d"

const FollowSetKey = "FollowKey"
const FollowTpl = "Follow&Follow&Id:%d"

const ExistDataValTpl = "%d:%d"
