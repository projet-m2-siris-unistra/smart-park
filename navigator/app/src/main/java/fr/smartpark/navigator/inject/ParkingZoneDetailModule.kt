package fr.smartpark.navigator.inject

import androidx.lifecycle.ViewModel
import dagger.Binds
import dagger.Module
import dagger.android.ContributesAndroidInjector
import dagger.multibindings.IntoMap
import fr.smartpark.navigator.ParkingZoneDetailFragment
import fr.smartpark.navigator.viewmodels.ParkingZoneDetailViewModel

@Module
abstract class ParkingZoneDetailModule {
    @ContributesAndroidInjector(modules = [
        ViewModelBuilder::class
    ])
    internal abstract fun parkingZoneDetailFragment(): ParkingZoneDetailFragment

    @Binds
    @IntoMap
    @ViewModelKey(ParkingZoneDetailViewModel::class)
    abstract fun bindViewModel(viewModel: ParkingZoneDetailViewModel): ViewModel
}