package fr.smartpark.navigator

import android.os.Bundle
import android.view.Menu
import android.view.MenuItem
import androidx.databinding.DataBindingUtil
import androidx.navigation.NavController
import androidx.navigation.NavDestination
import androidx.navigation.findNavController
import dagger.android.support.DaggerAppCompatActivity
import fr.smartpark.navigator.databinding.ActivityTenantsBinding

class TenantsActivity : DaggerAppCompatActivity(), NavController.OnDestinationChangedListener {
    private lateinit var binding: ActivityTenantsBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = DataBindingUtil.setContentView(this, R.layout.activity_tenants)
        setSupportActionBar(binding.tenantsToolbar)
        findNavController(R.id.nav_container).addOnDestinationChangedListener(this)
    }

    override fun onDestroy() {
        findNavController(R.id.nav_container).removeOnDestinationChangedListener(this)
        super.onDestroy()
    }

    override fun onDestinationChanged(
        controller: NavController,
        destination: NavDestination,
        arguments: Bundle?
    ) {
        when (destination.id) {
            R.id.settings ->
                supportActionBar?.title = "Settings"
            R.id.tenantList ->
                supportActionBar?.title = "SmartPark Navigator"
        }
    }
}
