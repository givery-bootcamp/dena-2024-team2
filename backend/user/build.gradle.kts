import org.jlleitschuh.gradle.ktlint.KtlintExtension

val kotlin_version: String by project
val logback_version: String by project
val exposed_version: String by project
val h2_version: String by project

plugins {
    kotlin("jvm") version "2.0.0"
    id("io.ktor.plugin") version "2.3.12"
    id("org.jetbrains.kotlin.plugin.serialization") version "2.0.0"
    id("com.google.devtools.ksp") version "2.0.0-1.0.22"
    id("org.jlleitschuh.gradle.ktlint") version "12.1.1"
    id("com.google.cloud.tools.jib") version "3.4.3"
}

group = "myapp"
version = "0.0.1"

application {
    mainClass.set("myapp.ApplicationKt")

    applicationDefaultJvmArgs = listOf("-Dio.ktor.development=true")
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(libs.bundles.ktor)
    implementation(libs.logback)
    implementation(libs.bundles.koin)
    implementation(libs.koin.ktor)
    implementation(libs.arrow)
    implementation(project(":repository"))
    testImplementation(libs.bundles.kotest)
    testImplementation(libs.test.kotlin)
    testImplementation(libs.test.koin)
    testImplementation(libs.test.ktor)
}

ksp {
    arg("KOIN_CONFIG_CHECK", "true")
}

jib {
    to {
        image = "101501319743.dkr.ecr.ap-northeast-1.amazonaws.com/dena-training-2024-team-2-backend-user"
        tags = setOf("latest")
    }
}

configure<KtlintExtension> {
    filter {
        exclude("build.gradle.kts")
        exclude("**/generated/**")
    }
}