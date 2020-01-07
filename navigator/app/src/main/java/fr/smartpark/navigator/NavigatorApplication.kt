package fr.smartpark.navigator

import com.google.android.libraries.maps.MapsInitializer
import dagger.android.AndroidInjector
import dagger.android.support.DaggerApplication
import fr.smartpark.navigator.inject.DaggerAppComponent

class NavigatorApplication : DaggerApplication() {
    override fun applicationInjector(): AndroidInjector<out DaggerApplication> {
        return DaggerAppComponent.factory().create(applicationContext)
    }

    override fun onCreate() {
        super.onCreate()
        MapsInitializer.initialize(this)
    }
}