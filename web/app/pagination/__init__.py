from sanic import Sanic, response
from sanic.exceptions import ServerError 
from app import ServerError
from math import ceil


class Pagination:
    
    """
    A class that creates an abstraction to handle pagination.
    It will allow to make easier calculation.
    Works only with Sanic
    """
    
    """
    You can personnalize page_size_choices which are the implemented page sizes
    """
    def __init__(self, request, page_size_choices = [20, 30, 40, 50]):
        self.page_size = 20     # default page size 
        self.page_number = 1    # default page number
        self.page_size_choices = page_size_choices
        self.computePageSize(request)
        self.computePageNumber(request)


    def computePageSize(self, request):
        if "pagesize" in request.raw_args:
            page_size = int(request.raw_args['pagesize'])
            
            if page_size not in self.page_size_choices:
                print("page_size ", page_size, " is not valid")
                raise ServerError("page size not valid", status_code=500)
            else:
                self.page_size = page_size


    def computePageNumber(self, request):
        if "page" in request.raw_args:
            page_number = int(request.raw_args['page'])
            if page_number < 1: # or page_number > page_max ==> to test
                print("page_number ", page_number, " is not valid")
                raise ServerError("page number not valid", status_code=500)
            else:
                self.page_number = page_number


    def setElementsNumber(self, number):
        # Number of total elements that can be shown
        self.elements_number = number
        # Maximum number of pages (I.e. the last page number)
        self.max_pages = ceil(self.elements_number / self.page_size)
