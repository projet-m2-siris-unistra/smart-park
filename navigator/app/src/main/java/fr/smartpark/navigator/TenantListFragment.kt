package fr.smartpark.navigator

import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.viewModels
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.observe
import androidx.navigation.fragment.findNavController
import dagger.android.support.DaggerFragment
import fr.smartpark.navigator.adapters.TenantAdapter
import fr.smartpark.navigator.api.ApiResult
import fr.smartpark.navigator.databinding.FragmentTenantListBinding
import fr.smartpark.navigator.viewmodels.TenantListViewModel
import javax.inject.Inject

class TenantListFragment : DaggerFragment() {
    @Inject
    lateinit var viewModelFactory: ViewModelProvider.Factory
    private val viewModel: TenantListViewModel by viewModels { viewModelFactory }

    private lateinit var binding: FragmentTenantListBinding

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentTenantListBinding.inflate(inflater, container, false)

        val adapter = TenantAdapter()
        binding.tenantList.adapter = adapter
        viewModel.tenants.observe(viewLifecycleOwner) {
            if (it.status == ApiResult.Status.SUCCESS) {
                adapter.submitList(it.data)
            }
        }
        return binding.root
    }
}