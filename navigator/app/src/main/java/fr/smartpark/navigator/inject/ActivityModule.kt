package fr.smartpark.navigator.inject

import dagger.Module
import dagger.android.ContributesAndroidInjector
import fr.smartpark.navigator.TenantsActivity
import fr.smartpark.navigator.ZonesActivity

@Module
abstract class ActivityModule {
    @ContributesAndroidInjector
    abstract fun contributeZonesActivity(): ZonesActivity

    @ContributesAndroidInjector
    abstract fun contributeTenantsActivity(): TenantsActivity
}