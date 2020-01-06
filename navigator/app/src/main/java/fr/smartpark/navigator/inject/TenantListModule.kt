package fr.smartpark.navigator.inject

import dagger.Module
import dagger.android.ContributesAndroidInjector
import fr.smartpark.navigator.TenantListFragment

@Module
abstract class TenantListModule {
    @ContributesAndroidInjector()
    internal abstract fun tenantListFragment(): TenantListFragment
}