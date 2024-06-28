import org.jlleitschuh.gradle.ktlint.KtlintExtension

plugins {
    alias(libs.plugins.kotlin)
    alias(libs.plugins.ktlint)
    alias(libs.plugins.ksp)
}

group = "myapp"
version = "0.0.1"

repositories {
    mavenCentral()
}

dependencies {
    implementation(libs.bundles.koin)
    ksp(libs.koin.ksp)
    api(project(":entity"))
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
