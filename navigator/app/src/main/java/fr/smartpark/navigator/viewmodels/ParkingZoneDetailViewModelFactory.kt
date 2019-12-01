package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import fr.smartpark.navigator.data.ParkingZoneRepository

class ParkingZoneDetailViewModelFactory(
    private val zoneRepository: ParkingZoneRepository,
    private val zoneId: String
) : ViewModelProvider.NewInstanceFactory() {
    @Suppress("UNCHECKED_CAST")
    override fun <T : ViewModel?> create(modelClass: Class<T>) =
        ParkingZoneDetailViewModel(zoneRepository, zoneId) as T
}
