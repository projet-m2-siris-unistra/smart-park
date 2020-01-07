package fr.smartpark.navigator

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.viewModels
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.observe
import dagger.android.support.DaggerFragment
import fr.smartpark.navigator.adapters.ParkingZoneAdapter
import fr.smartpark.navigator.api.ApiResult
import fr.smartpark.navigator.databinding.FragmentParkingZoneListBinding
import fr.smartpark.navigator.viewmodels.ParkingZoneListViewModel
import javax.inject.Inject

class ParkingZoneListFragment : DaggerFragment() {
    @Inject
    lateinit var viewModelFactory: ViewModelProvider.Factory
    private val viewModel by viewModels<ParkingZoneListViewModel> { viewModelFactory }

    private lateinit var binding: FragmentParkingZoneListBinding

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentParkingZoneListBinding.inflate(inflater, container, false)

        val tenantId = (requireActivity() as ZonesActivity).getTenantId()
        viewModel.start(tenantId)

        val adapter = ParkingZoneAdapter()
        binding.zoneList.adapter = adapter
        viewModel.zones.observe(viewLifecycleOwner) { zones ->
            if (zones.status == ApiResult.Status.SUCCESS) {
                binding.zoneList.visibility = View.VISIBLE
                binding.progressBar.visibility = View.GONE
                adapter.submitList(zones.data!!)
            }
        }
        return binding.root
    }
}
