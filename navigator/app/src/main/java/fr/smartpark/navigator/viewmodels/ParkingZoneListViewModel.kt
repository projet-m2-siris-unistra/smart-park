package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.switchMap
import fr.smartpark.navigator.api.ApiResult
import fr.smartpark.navigator.data.models.Zone
import fr.smartpark.navigator.data.ZoneRepository
import javax.inject.Inject

class ParkingZoneListViewModel @Inject constructor(private val zoneRepository: ZoneRepository) :
    ViewModel() {
    private val _tenantId = MutableLiveData<Long>()
    val zones: LiveData<ApiResult<List<Zone>>> = _tenantId.switchMap { zoneRepository.getZones(it) }
    fun start(tenantId: Long) {
        _tenantId.postValue(tenantId)
    }
}