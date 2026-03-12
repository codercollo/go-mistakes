package internal

//Simulate redis package import alias to avoid name collision
//In real code: import redisapi "mylib/redis"

//BAD: variable shadows the package name:
// redis := redis.NewClient  //redis package now inaccessible in this scope

// v, err := redis.Get("foo")

//GOOD: option 1 - rename the variable:
//				redisClient := redis.NewClient() // clear no collision

//GOOD option2 - alias the import:
//			import redisapi "mylib/redis"
//			redis := redisapi.NewClient() //qualifier is unambiguous

//Also avoid shadowing built-ins
// 			copy := copyFile(src, dst) //built in copy() now hiddens
