package myapp.model

data class UserLogin(
    val name: String,
    val password: String,
)

data class UserPublic(
    val id: Int,
    val name: String,
)
