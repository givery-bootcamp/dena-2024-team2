package myapp.usecase

import org.koin.dsl.bind
import org.koin.dsl.module

val usecaseModule = module {
    single { UserSigninUsecaseImpl(get()) }.bind(UserSigninUsecase::class)
    single { NewUserUsecaseImpl(get()) }.bind(NewUserUsecase::class)
}
