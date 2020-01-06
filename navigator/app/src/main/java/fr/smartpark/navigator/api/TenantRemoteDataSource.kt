package fr.smartpark.navigator.api

import javax.inject.Inject

class TenantRemoteDataSource @Inject constructor(private val service: ApiService) : BaseDataSource() {
    suspend fun getTenants() = getResult { service.listTenants() }.map { it.tenants }
}
