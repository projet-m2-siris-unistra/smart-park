package fr.smartpark.navigator.api

import javax.inject.Inject

class ZoneRemoteDataSource @Inject constructor(private val service: ApiService) : BaseDataSource() {
    suspend fun getZones() = getResult { service.listZones() }
}