package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.*
import fr.smartpark.navigator.data.models.Zone
import fr.smartpark.navigator.data.ZoneRepository
import javax.inject.Inject

class ParkingZoneDetailViewModel @Inject constructor(
    private val zoneRepository: ZoneRepository
) : ViewModel() {
    private val _params = MutableLiveData<Pair<Long, Long>>()
    val zone: LiveData<Zone> = _params.switchMap { zoneRepository.getZone(it.first, it.second) }

    fun start(tenantId: Long, zoneId: Long) {
        _params.postValue(Pair(tenantId, zoneId))
    }
}