package fr.smartpark.navigator.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import fr.smartpark.navigator.api.ApiResult
import fr.smartpark.navigator.data.TenantRepository
import fr.smartpark.navigator.data.models.Tenant
import javax.inject.Inject

class TenantListViewModel @Inject constructor(tenantRepository: TenantRepository) :
    ViewModel() {
    val tenants: LiveData<ApiResult<List<Tenant>>> = tenantRepository.tenants
}
