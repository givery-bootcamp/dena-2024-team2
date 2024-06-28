package myapp

data class User(
    val id: Int,
    val name: String,
    val password: String,
)

data class UserLogin(
    val name: String,
    val password: String,
)

data class UserPublic(
    val id: Int,
    val name: String,
)
