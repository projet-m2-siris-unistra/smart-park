package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import fr.smartpark.navigator.data.ParkingZoneRepository

class ParkingZoneListViewModelFactory (
    private val repository: ParkingZoneRepository
) : ViewModelProvider.NewInstanceFactory() {
    @Suppress("UNCHECKED_CAST")
    override fun <T : ViewModel?> create(modelClass: Class<T>) =
        ParkingZoneListViewModel(repository) as T
}