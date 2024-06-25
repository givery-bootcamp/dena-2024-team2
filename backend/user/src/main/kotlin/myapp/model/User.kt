package myapp.model

data class User(
    val id: UInt,
    val name: String,
    val password: String
)

data class UserLogin(
    val name: String,
    val password: String
)

data class UserPublic(
    val id: UInt,
    val name: String
)