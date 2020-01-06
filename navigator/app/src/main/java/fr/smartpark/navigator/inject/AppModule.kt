package fr.smartpark.navigator.inject

import android.content.Context
import androidx.room.Room
import dagger.Module
import dagger.Provides
import fr.smartpark.navigator.api.ApiService
import fr.smartpark.navigator.api.TenantRemoteDataSource
import fr.smartpark.navigator.api.ZoneRemoteDataSource
import fr.smartpark.navigator.data.*
import fr.smartpark.navigator.utilities.DATABASE_NAME
import kotlinx.coroutines.Dispatchers
import javax.inject.Singleton

@Module
object AppModule {
    @Singleton
    @Provides
    fun provideTenantRemoteDataSource(service: ApiService) =
        TenantRemoteDataSource(service)

    @Singleton
    @Provides
    fun provideTenantDao(database: AppDatabase) =
        database.tenantDao()

    @Singleton
    @Provides
    fun provideTenantRepository(dao: TenantDao, remoteDataSource: TenantRemoteDataSource) =
        TenantRepository(dao, remoteDataSource)

    @Singleton
    @Provides
    fun provideZoneRemoteDataSource(service: ApiService) =
        ZoneRemoteDataSource(service)

    @Singleton
    @Provides
    fun provideZoneDao(database: AppDatabase) =
        database.zoneDao()

    @Singleton
    @Provides
    fun provideZoneRepository(dao: ZoneDao, remoteDataSource: ZoneRemoteDataSource) =
        ZoneRepository(dao, remoteDataSource)

    @Singleton
    @Provides
    fun provideDatabase(context: Context) =
        Room.databaseBuilder(context.applicationContext, AppDatabase::class.java, DATABASE_NAME)
            .fallbackToDestructiveMigration()
            .build()

    @Singleton
    @Provides
    fun provideIoDispatcher() = Dispatchers.IO
}
