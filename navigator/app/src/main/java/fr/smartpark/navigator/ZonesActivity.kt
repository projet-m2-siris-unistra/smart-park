package fr.smartpark.navigator

import android.content.Intent
import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.databinding.DataBindingUtil
import androidx.lifecycle.observe
import androidx.navigation.navArgs
import com.google.android.libraries.maps.CameraUpdateFactory
import com.google.android.libraries.maps.GoogleMap
import com.google.android.libraries.maps.OnMapReadyCallback
import com.google.android.libraries.maps.SupportMapFragment
import com.google.android.libraries.maps.model.LatLng
import com.google.android.libraries.maps.model.MarkerOptions
import dagger.android.AndroidInjection
import dagger.android.DaggerActivity
import dagger.android.HasAndroidInjector
import dagger.android.support.DaggerAppCompatActivity
import fr.smartpark.navigator.data.TenantRepository
import fr.smartpark.navigator.databinding.ActivityZonesBinding
import javax.inject.Inject

class ZonesActivity : DaggerAppCompatActivity(), OnMapReadyCallback {
    private lateinit var binding: ActivityZonesBinding
    private val args: ZonesActivityArgs by navArgs()
    private var _tenantId: Long? = null

    @Inject
    lateinit var tenantRepository: TenantRepository

    fun getTenantId(): Long = when(_tenantId) {
        null -> args.tenantId
        else -> _tenantId!!
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = DataBindingUtil.setContentView(this, R.layout.activity_zones)
        setSupportActionBar(binding.tenantsToolbar)

        if (intent.action == Intent.ACTION_VIEW) {
            _tenantId = intent.data?.
                getQueryParameter("tenant_id")?.
                toLongOrNull()
        }

        val mapFragment: SupportMapFragment = supportFragmentManager.findFragmentById(R.id.map)!! as SupportMapFragment
        mapFragment.getMapAsync(this)
    }

    override fun onMapReady(map: GoogleMap?) {
        tenantRepository.getTenant(getTenantId()).observe(this) {
            supportActionBar?.title = it.name
            val pos = LatLng(it.geo.lat.toDouble(), it.geo.long.toDouble())
            map?.addMarker(MarkerOptions().position(pos))
            map?.moveCamera(CameraUpdateFactory.newLatLngZoom(pos, 10.0F))
        }
    }
}
