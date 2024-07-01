package myapp.repository

import org.jetbrains.exposed.sql.Table

internal object UserTable : Table("users") {
    val id = integer("id").autoIncrement()
    val name = varchar("name", 32).uniqueIndex()
    val password = char("password", 60)

    override val primaryKey: PrimaryKey = PrimaryKey(id)
}
