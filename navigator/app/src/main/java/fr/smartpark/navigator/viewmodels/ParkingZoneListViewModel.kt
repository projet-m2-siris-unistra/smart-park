package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import fr.smartpark.navigator.api.ApiResult
import fr.smartpark.navigator.data.models.Zone
import fr.smartpark.navigator.data.ZoneRepository
import javax.inject.Inject

class ParkingZoneListViewModel @Inject constructor(zoneRepository: ZoneRepository) :
    ViewModel() {
    val zones: LiveData<ApiResult<List<Zone>>> = zoneRepository.zones
}