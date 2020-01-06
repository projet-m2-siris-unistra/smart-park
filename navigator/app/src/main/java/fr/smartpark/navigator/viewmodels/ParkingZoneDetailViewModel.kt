package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.*
import fr.smartpark.navigator.data.models.Zone
import fr.smartpark.navigator.data.ZoneRepository
import javax.inject.Inject

class ParkingZoneDetailViewModel @Inject constructor(
    zoneRepository: ZoneRepository
) : ViewModel() {
    private val _zoneId = MutableLiveData<Long>()
    val zone: LiveData<Zone> = _zoneId.switchMap { id -> zoneRepository.getZone(id) }

    fun start(zoneId: Long) {
        _zoneId.postValue(zoneId)
    }
}