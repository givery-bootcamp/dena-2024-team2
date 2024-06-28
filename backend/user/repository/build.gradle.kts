import org.jlleitschuh.gradle.ktlint.KtlintExtension

plugins {
    alias(libs.plugins.kotlin)
    alias(libs.plugins.ktlint)
}

group = "myapp"
version = "0.0.1"

repositories {
    mavenCentral()
}

dependencies {
    implementation(libs.bundles.exposed)
    implementation(libs.bundles.koin)
    implementation(project(":usecase"))
    testImplementation(libs.test.kotlin)
    testImplementation(libs.bundles.kotest)
    testImplementation(libs.test.koin)
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(21)
}

configure<KtlintExtension> {
    filter {
        exclude("build.gradle.kts")
        exclude("**/generated/**")
    }
}
