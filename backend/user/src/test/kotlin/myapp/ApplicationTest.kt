package myapp

import io.kotest.core.spec.style.FunSpec
import io.kotest.matchers.shouldBe
import io.ktor.client.request.get
import io.ktor.client.statement.bodyAsText
import io.ktor.http.HttpStatusCode
import io.ktor.server.testing.testApplication
import myapp.plugins.configureRouting

class ApplicationTest : FunSpec({
    test("GET / should return text") {
        testApplication {
            application { configureRouting() }
            client.get("/").apply {
                status shouldBe HttpStatusCode.OK
                bodyAsText() shouldBe "Hello World!"
            }
        }
    }
})
