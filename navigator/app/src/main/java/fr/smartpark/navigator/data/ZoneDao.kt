package fr.smartpark.navigator.data

import androidx.lifecycle.LiveData
import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query
import fr.smartpark.navigator.data.models.Zone

@Dao
interface ZoneDao {
    @Query("SELECT * FROM zones WHERE tenant_id = :tenantId ORDER BY id")
    fun getZones(tenantId: Long): LiveData<List<Zone>>

    @Query("SELECT * from zones WHERE tenant_id = :tenantId AND id = :id")
    fun getZone(tenantId: Long, id: Long): LiveData<Zone>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertAll(zones: List<Zone>)
}