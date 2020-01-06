package fr.smartpark.navigator.inject

import androidx.lifecycle.ViewModel
import dagger.Binds
import dagger.Module
import dagger.android.ContributesAndroidInjector
import dagger.multibindings.IntoMap
import fr.smartpark.navigator.ParkingZoneListFragment
import fr.smartpark.navigator.viewmodels.ParkingZoneListViewModel

@Module
abstract class ParkingZoneListModule {
    @ContributesAndroidInjector(modules = [
        ViewModelBuilder::class
    ])
    internal abstract fun parkingZoneListFragment(): ParkingZoneListFragment

    @Binds
    @IntoMap
    @ViewModelKey(ParkingZoneListViewModel::class)
    abstract fun bindViewModel(viewModel: ParkingZoneListViewModel): ViewModel
}
