import { create } from "zustand";




const useBrandCategoryStore = create ((set) => ({
    menu_id: '',
    changeMenu_id: ((id:string) => {
        set({ menu_id: id });
    })
}));

export default useBrandCategoryStore;