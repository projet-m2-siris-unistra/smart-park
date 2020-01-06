package fr.smartpark.navigator.data.models

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey
import com.google.gson.annotations.SerializedName

@Entity(tableName = "tenants")
data class Tenant(
    @PrimaryKey @ColumnInfo(name = "id") @SerializedName("tenant_id") val tenantId: Long,
    val name: String
)

