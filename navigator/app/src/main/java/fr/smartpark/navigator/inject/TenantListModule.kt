package fr.smartpark.navigator.inject

import androidx.lifecycle.ViewModel
import dagger.Binds
import dagger.Module
import dagger.android.ContributesAndroidInjector
import dagger.multibindings.IntoMap
import fr.smartpark.navigator.TenantListFragment
import fr.smartpark.navigator.viewmodels.TenantListViewModel

@Module
abstract class TenantListModule {
    @ContributesAndroidInjector(modules = [
        ViewModelBuilder::class
    ])
    internal abstract fun tenantListFragment(): TenantListFragment

    @Binds
    @IntoMap
    @ViewModelKey(TenantListViewModel::class)
    abstract fun bindViewModel(viewModel: TenantListViewModel): ViewModel
}