package fr.smartpark.navigator.inject

import android.content.Context
import dagger.BindsInstance
import dagger.Component
import dagger.android.AndroidInjector
import dagger.android.support.AndroidSupportInjectionModule
import fr.smartpark.navigator.NavigatorApplication
import javax.inject.Singleton

@Singleton
@Component(modules = [
    AppModule::class,
    AndroidSupportInjectionModule::class,
    ApiModule::class,
    ParkingZoneDetailModule::class,
    ParkingZoneListModule::class,
    TenantListModule::class
])
interface AppComponent : AndroidInjector<NavigatorApplication> {
    @Component.Factory
    interface Factory {
        fun create(@BindsInstance applicationContext: Context): AppComponent
    }
}
