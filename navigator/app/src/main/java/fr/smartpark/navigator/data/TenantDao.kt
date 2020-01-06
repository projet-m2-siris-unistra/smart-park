package fr.smartpark.navigator.data

import androidx.lifecycle.LiveData
import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query
import fr.smartpark.navigator.data.models.Tenant

@Dao
interface TenantDao {
    @Query("SELECT * FROM tenants ORDER BY id")
    fun getTenants(): LiveData<List<Tenant>>

    @Query("SELECT * from tenants WHERE id = :id")
    fun getTenant(id: Long): LiveData<Tenant>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertAll(tenants: List<Tenant>)
}
